// this file was generated by protoc-gen-gotemplate

package entity

import (
	"github.com/pkg/errors"

	"berty.tech/core/pkg/errorcodes"
	"berty.tech/core/pkg/validate/validator"
)

var (
	_ = errors.New
	_ = validator.IsContactKey
	_ = errorcodes.IsSubCode
)

func (m *Event) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: ID - name:"id" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"id" options:<53004:1 []:"ID" 65006:"gorm:\"primary_key\"" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: SourceDeviceID - name:"source_device_id" number:2 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"sourceDeviceId" options:<[]:"SourceDeviceID" 65006:"gorm:\"primary_key\"" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: CreatedAt - name:"created_at" number:3 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"createdAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: CreatedAt")
		}
	}

	// handling field: UpdatedAt - name:"updated_at" number:4 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"updatedAt" options:<65001:0 65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: UpdatedAt")
		}
	}

	// handling field: SentAt - name:"sent_at" number:6 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"sentAt" options:<65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetSentAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: SentAt")
		}
	}

	// handling field: ReceivedAt - name:"received_at" number:7 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"receivedAt" options:<65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetReceivedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: ReceivedAt")
		}
	}

	// handling field: AckedAt - name:"acked_at" number:8 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"ackedAt" options:<65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetAckedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: AckedAt")
		}
	}

	// handling field: Direction - name:"direction" number:9 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.entity.Event.Direction" json_name:"direction" options:<[]:0 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: APIVersion - name:"api_version" number:10 label:LABEL_OPTIONAL type:TYPE_UINT32 json_name:"apiVersion" options:<[]:"APIVersion" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Kind - name:"kind" number:13 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.entity.Kind" json_name:"kind" options:<[]:0 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Attributes - name:"attributes" number:14 label:LABEL_OPTIONAL type:TYPE_BYTES json_name:"attributes" options:<53005:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: SeenAt - name:"seen_at" number:16 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"seenAt" options:<65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetSeenAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: SeenAt")
		}
	}

	// handling field: AckStatus - name:"ack_status" number:17 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.entity.Event.AckStatus" json_name:"ackStatus" options:<[]:0 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: Dispatches - name:"dispatches" number:18 label:LABEL_REPEATED type:TYPE_MESSAGE type_name:".berty.entity.EventDispatch" json_name:"dispatches"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetDispatches()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Dispatches")
		}
	}

	// handling field: SourceContactID - name:"source_contact_id" number:19 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"sourceContactId" options:<[]:"SourceContactID" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: TargetType - name:"target_type" number:20 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.entity.Event.TargetType" json_name:"targetType"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: TargetAddr - name:"target_addr" number:21 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"targetAddr" options:<53004:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: ErrProxy - name:"err_proxy" number:98 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".berty.entity.Err" json_name:"errProxy" options:<65006:"gorm:\"-\"" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetErrProxy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: ErrProxy")
		}
	}

	// handling field: Metadata - name:"metadata" number:99 label:LABEL_REPEATED type:TYPE_MESSAGE type_name:".berty.entity.MetadataKeyValue" json_name:"metadata" options:<65006:"gorm:\"-\"" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: Metadata")
		}
	}

	return nil
}
func (m *EventDispatch) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: EventID - name:"event_id" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"eventId" options:<[]:"EventID" 65006:"gorm:\"primary_key\"" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: DeviceID - name:"device_id" number:2 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"deviceId" options:<[]:"DeviceID" 65006:"gorm:\"primary_key\"" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: ContactID - name:"contact_id" number:3 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"contactId" options:<[]:"ContactID" >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: SentAt - name:"sent_at" number:4 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"sentAt" options:<65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetSentAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: SentAt")
		}
	}

	// handling field: AckedAt - name:"acked_at" number:5 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"ackedAt" options:<65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetAckedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: AckedAt")
		}
	}

	// handling field: SeenAt - name:"seen_at" number:6 label:LABEL_OPTIONAL type:TYPE_MESSAGE type_name:".google.protobuf.Timestamp" json_name:"seenAt" options:<65010:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	if v, ok := interface{}(m.GetSeenAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return errors.Wrap(err, "embedded message verification failed: SeenAt")
		}
	}

	// handling field: AckMedium - name:"ack_medium" number:7 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.entity.EventDispatch.Medium" json_name:"ackMedium"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: SeenMedium - name:"seen_medium" number:8 label:LABEL_OPTIONAL type:TYPE_ENUM type_name:".berty.entity.EventDispatch.Medium" json_name:"seenMedium"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)

	// handling field: RetryBackoff - name:"retry_backoff" number:9 label:LABEL_OPTIONAL type:TYPE_INT64 json_name:"retryBackoff"  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	return nil
}
func (m *MetadataKeyValue) Validate() error {
	if m == nil {
		return nil
	}

	// handling field: Key - name:"key" number:1 label:LABEL_OPTIONAL type:TYPE_STRING json_name:"key" options:<[]:1 >  (is_contact_key=false, defined_only=false, min_len=1, max_len=0, skip=false, required=false, min_items=0, max_items=0)
	if len(m.GetKey()) < 1 {
		return errors.New("Key must be longer than 1")
	}

	// handling field: Values - name:"values" number:2 label:LABEL_REPEATED type:TYPE_STRING json_name:"values" options:<[]:1 >  (is_contact_key=false, defined_only=false, min_len=0, max_len=0, skip=false, required=false, min_items=1, max_items=0)
	if len(m.GetValues()) < 1 {
		return errors.New("Values must contain at least 1 item(s)")
	}
	return nil
}
